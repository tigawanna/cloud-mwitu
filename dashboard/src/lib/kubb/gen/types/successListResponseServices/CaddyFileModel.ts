/**
 * @description SuccessListResponse_services.CaddyFileModel schema
 */
export type SuccessListResponseServicesCaddyFileModel = {
  error?: unknown | undefined
  /**
   * @type object | undefined
   */
  result?:
    | {
        /**
         * @type array | undefined
         */
        items?:
          | {
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
            }[]
          | undefined
      }
    | undefined
}