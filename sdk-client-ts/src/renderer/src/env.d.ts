/// <reference types="vite/client" />
declare module '*.vue' {
  import { CommonExecOptions } from 'vue'

  const commonExecOptions: CommonExecOptions
  export default commonExecOptions
}

declare module '*.js' {
  import { CommonExecOptions } from 'vue'

  const commonExecOptions: CommonExecOptions
  export default commonExecOptions
}
