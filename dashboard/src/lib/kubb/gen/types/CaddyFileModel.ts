/**
 * @description CaddyFileModel schema
 */
export type CaddyFileModel = {
  /**
   * @type array | undefined
   */
  block?:
    | {
        /**
         * @type array | undefined
         */
        path?: string[] | undefined
        /**
         * @type array | undefined
         */
        value?: string[] | undefined
      }[]
    | undefined
  /**
   * @type string | undefined
   */
  content?: string | undefined
  /**
   * @type string | undefined
   */
  domain?: string | undefined
  startEnd?: unknown | undefined
}