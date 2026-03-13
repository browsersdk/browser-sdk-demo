/* eslint-disable */
import { app } from 'electron'
import path from 'path'
import koffi from 'koffi'

/* ============================================================
   1. 加载 brosdk.dll
   ============================================================ */

export default class BroSDK {
  _DLL_PATH
  #lib
  /* ============================================================
   2. 类型定义
   ============================================================ */
  // SDK 回调原型: void __cdecl (int32_t code, void *ud, const char *data, size_t len)
  // registerAsync 确保回调在 Node.js 主线程上触发（线程安全）
  #ResultCbProto
  // char* 和 char** 指针类型，用于 sdk_init 的输出参数
  #CharPtr
  #CharPtrPtr
  /* ============================================================
   3. 函数绑定
   ============================================================ */
  $_register_cb
  // sdk_init 有两个输出参数: char **out_data, size_t *out_len
  // SDK 句柄输出槽：sdk_handle_t* = void**，x64 上 size_t 与指针同宽，用作不透明槽
  #HandleSlot
  $_init
  $_init_async
  $_init_webapi
  $_browser_open
  $_browser_close
  $_token_update
  $_shutdown
  $_free
  $_error_name
  $_error_string
  $_is_error
  $_is_warn
  $_is_reqid
  $_is_ok
  $_is_done
  /* ============================================================
   4. 回调注册
   ============================================================ */
  /** JS 层用户回调 */
  #resultCb: Function | null = null
  /** koffi 持有的 C-callable 函数指针，防止 GC */
  #nativeCb: koffi.IKoffiRegisteredCallback | null = null

  constructor() {
    console.log('app.isPackaged', app.isPackaged, app.getPath('exe'), app.getPath('home'))
    this._DLL_PATH = !app.isPackaged
      ? path.join(app.getAppPath(), 'sdk', 'brosdk.dll')
      : path.join(process.resourcesPath, 'sdk', 'brosdk.dll')
    // this._DLL_PATH = path.join(__dirname, "brosdk", "brosdk.dll");
    this.#lib = koffi.load(this._DLL_PATH)

    this.#ResultCbProto = koffi.proto('void __cdecl sdk_result_cb_t(int32_t, void *, const char *, size_t)')

    this.#CharPtr = koffi.pointer('uint8')
    this.#CharPtrPtr = koffi.pointer(this.#CharPtr)

    this.$_register_cb = this.#lib.func('int32_t __cdecl sdk_register_result_cb(sdk_result_cb_t *cb, void *ud)')
    this.#HandleSlot = koffi.out(koffi.pointer('size_t'))

    this.$_init = this.#lib.func('sdk_init', 'int32', [
      this.#HandleSlot, // sdk_handle_t * (不使用多实例，仅提供写入槽)
      'str', // const char *data
      'size_t', // size_t len
      koffi.out(this.#CharPtrPtr), // char **out_data  —— SDK 分配，调用方 sdk_free
      koffi.out(koffi.pointer('size_t')) // size_t *out_len
    ])
    this.$_init_async = this.#lib.func('sdk_init_async', 'int32', [this.#HandleSlot, 'str', 'size_t'])

    this.$_init_webapi = this.#lib.func('int32_t __cdecl sdk_init_webapi(uint16_t port)')
    this.$_browser_open = this.#lib.func('int32_t __cdecl sdk_browser_open(const char *data, size_t len)')
    this.$_browser_close = this.#lib.func('int32_t __cdecl sdk_browser_close(const char *data, size_t len)')
    this.$_token_update = this.#lib.func('int32_t __cdecl sdk_token_update(const char *data, size_t len)')
    this.$_shutdown = this.#lib.func('int32_t __cdecl sdk_shutdown(void)')
    this.$_free = this.#lib.func('void __cdecl sdk_free(void *ptr)')
    this.$_error_name = this.#lib.func('const char * __cdecl sdk_error_name(int32_t code)')
    this.$_error_string = this.#lib.func('const char * __cdecl sdk_error_string(int32_t code)')
    this.$_is_error = this.#lib.func('bool __cdecl sdk_is_error(int32_t code)')
    this.$_is_warn = this.#lib.func('bool __cdecl sdk_is_warn(int32_t code)')
    this.$_is_reqid = this.#lib.func('bool __cdecl sdk_is_reqid(int32_t code)')
    this.$_is_ok = this.#lib.func('bool __cdecl sdk_is_ok(int32_t code)')
    this.$_is_done = this.#lib.func('bool __cdecl sdk_is_done(int32_t code)')
  }
  registerResultCb = (fn: () => void | null) => {
    this.#resultCb = fn
    // 释放旧的回调句柄
    if (this.#nativeCb) {
      koffi.unregister(this.#nativeCb)
      this.#nativeCb = null
    }
    console.log('this.#ResultCbProto', this.#ResultCbProto)
    // register: SDK 从内部线程调用此指针时，koffi 将调用调度到 Node.js 主线程
    this.#nativeCb = koffi.register(
      (code, _ud, data, _len) => {
        if (this.#resultCb) this.#resultCb(code, data || '')
      },
      koffi.pointer(this.#ResultCbProto) // koffi.register 需要 pointer-to-proto
    )
    return this.$_register_cb(this.#nativeCb, null)
  }
  // Accepts a string or a plain object (auto-serialised to JSON).
  // All DLL entry points pass their payload through here to get the byte length.
  strLen = (s) => {
    return Buffer.byteLength(s, 'utf8')
  }
  // Normalise a caller-supplied payload to a UTF-8 JSON string.
  // Accepts: string | object | null | undefined.
  #toJson = (v) => {
    if (v === null || v === undefined) return '{}'
    if (typeof v === 'string') return v
    return JSON.stringify(v)
  }
  init = (json) => {
    const input = this.#toJson(json)
    const handle = [0]
    let outData = [null]
    let outLen = [0]
    const ret = this.$_init(handle, input, this.strLen(input), outData, outLen)
    const len = Number(outLen[0])
    let response: string | null = null
    if (this.$_is_ok(ret) && outData[0] && len > 0) {
      const bytes = koffi.decode(outData[0], 'uint8', len)
      response = Buffer.from(bytes).toString('utf8')
      // DO NOT free outData[0] here — caller is responsible for freePointer(ptr)
    }
    return { code: ret, ptr: outData[0], len, response }
  }

  /**
   * 释放由 `init` 返回的 ptr
   */
  freePointer = (ptr) => {
    if (ptr) this.$_free(ptr)
  }
  /** 异步初始化，结果通过 registerResultCb 回调返回。 */
  initAsync = (json) => {
    const input = this.#toJson(json)
    const handle = [0]
    return this.$_init_async(handle, input, this.strLen(input))
  }

  initWebAPI = (port) => {
    return this.$_init_webapi(port)
  }
  browserOpen = (json) => {
    const s = this.#toJson(json)
    return this.$_browser_open(s, this.strLen(s))
  }
  browserClose = (json) => {
    const s = this.#toJson(json)
    return this.$_browser_close(s, this.strLen(s))
  }
  tokenUpdate = (json) => {
    const s = this.#toJson(json)
    return this.$_token_update(s, this.strLen(s))
  }

  shutdown = () => {
    const ret = this.$_shutdown()
    if (this.#nativeCb) {
      koffi.unregister(this.#nativeCb)
      this.#nativeCb = null
    }
    return ret
  }

  errorName = (code) => {
    return this.$_error_name(code) || ''
  }
  errorString = (code) => {
    return this.$_error_string(code) || ''
  }
  isOk = (code) => {
    console.log('isOk', code)
    return this.$_is_ok(code)
  }
  isError = (code) => {
    return this.$_is_error(code)
  }
  isDone = (code) => {
    return this.$_is_done(code)
  }
  isReqid = (code) => {
    return this.$_is_reqid(code)
  }
  isWarn = (code) => {
    return this.$_is_warn(code)
  }
  /**
   * Read cookies for a browser environment.
   * @param {bigint|number|string} envId  - The numeric environment ID (uint64).
   * @returns {{ code: number, ptr: External|null, len: number, response: string|null }}
   *   `response` is the raw JSON array string "[{...}, ...]".  Call freePointer(ptr) when done.
   */
  // readCookies = (envId) => {
  //   // Build JSON manually to avoid JSON.stringify's BigInt limitation
  //   const input = `{"envId":${envId}}`
  //   let outData = [null]
  //   let outLen = [0]
  //   const ret = this.$_read_cookies(input, this.strLen(input), outData, outLen)
  //   const len = Number(outLen[0])
  //   let response: string | null = null
  //   if (this.$_is_ok(ret) && outData[0] && len > 0) {
  //     const bytes = koffi.decode(outData[0], 'uint8', len)
  //     response = Buffer.from(bytes).toString('utf8')
  //   }
  //   return { code: ret, ptr: outData[0], len, response }
  // }

  printErrno = (tag, code) => {
    const label = this.isOk(code) ? 'OK' : this.isDone(code) ? 'DONE' : this.isError(code) ? 'ERROR' : 'UNKNOWN'
    const name = this.errorName(code)
    const msg = this.errorString(code)
    console.log(`[${tag}] ${label}  code=${code}  (${name}): ${msg}`)
  }
}
