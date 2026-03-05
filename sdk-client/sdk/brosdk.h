#if !defined(__BROSDK_H__)
#define __BROSDK_H__
/* Minimal public C API header - keep footprint small but include C types
  needed by language bindings (cgo, ctypes, etc.). */
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#if defined(SDK_API)
#undef SDK_API
#endif
#if defined(SDK_CALL)
#undef SDK_CALL
#endif

#if defined(_WIN32) || defined(_WIN64)
#define SDK_API __declspec(dllexport)
#define SDK_CALL __cdecl
#elif defined(__clang__) || defined(__GNUC__)
#define SDK_API __attribute__((visibility("default")))
#define SDK_CALL
#else
#define SDK_API
#define SDK_CALL
#endif

#define SDK_WARN_BASE 0x64
#define SDK_REQID_BASE 0xFF

#ifdef __cplusplus
extern "C" {
#endif
using sdk_errno_t = int32_t;
using sdk_port_t = uint16_t;
typedef void *sdk_handle_t;
typedef void(SDK_CALL *sdk_result_cb_t)(int32_t code, void *user_data,
                                        const char *data, size_t len);

SDK_API int32_t SDK_CALL sdk_register_result_cb(sdk_result_cb_t cb,
                                                void *user_data);
SDK_API int32_t SDK_CALL sdk_init_webapi(uint16_t port);
SDK_API int32_t SDK_CALL sdk_init_async(sdk_handle_t *cpp_handle,
                                        const char *data, size_t len);
SDK_API int32_t SDK_CALL sdk_init(sdk_handle_t *cpp_handle, const char *data,
                                  size_t len, char **out_data, size_t *out_len);
SDK_API int32_t SDK_CALL sdk_browser_open(const char *data, size_t len);
SDK_API int32_t SDK_CALL sdk_browser_close(const char *data, size_t len);
SDK_API int32_t SDK_CALL sdk_token_update(const char *data, size_t len);
SDK_API int32_t SDK_CALL sdk_shutdown(void);
SDK_API void SDK_CALL sdk_free(void *ptr);
SDK_API const char *SDK_CALL sdk_error_name(int32_t code);
SDK_API const char *SDK_CALL sdk_error_string(int32_t code);
SDK_API bool SDK_CALL sdk_is_error(int32_t code);
SDK_API bool SDK_CALL sdk_is_warn(int32_t code);
SDK_API bool SDK_CALL sdk_is_reqid(int32_t code);
SDK_API bool SDK_CALL sdk_is_ok(int32_t code);
SDK_API bool SDK_CALL sdk_is_done(int32_t code);

#ifdef __cplusplus
}
#endif

/* C++ interface (kept in the same header but only visible to C++ callers) */
#ifdef __cplusplus
class ISDK {
public:
  virtual ~ISDK() = default;
  virtual int32_t InitAsync(const char *data, size_t len) = 0;
  virtual int32_t Init(const char *data, size_t len, char **out_data,
                       size_t *out_len) = 0;
  virtual int32_t InitWebAPI(uint16_t port) = 0;
  virtual int32_t BrowserOpen(const char *data, size_t len) = 0;
  virtual int32_t BrowserClose(const char *data, size_t len) = 0;
  virtual int32_t RegisterResultCb(sdk_result_cb_t cb, void *user_data) = 0;
  virtual int32_t Shutdown() = 0;
};
#endif

#endif ///__BROSDK_H__