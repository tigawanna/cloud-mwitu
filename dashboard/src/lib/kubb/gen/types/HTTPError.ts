/**
 * @description HTTPError schema
 */
export type HTTPError = {
  /**
   * @description Human readable error message
   * @type string
   */
  detail?: (string | null) | undefined
  /**
   * @type array
   */
  errors?:
    | (
        | {
            /**
             * @type object | undefined
             */
            more?:
              | {
                  [key: string]: unknown
                }
              | undefined
            /**
             * @type string | undefined
             */
            name?: string | undefined
            /**
             * @type string | undefined
             */
            reason?: string | undefined
          }[]
        | null
      )
    | undefined
  /**
   * @type string
   */
  instance?: (string | null) | undefined
  /**
   * @description HTTP status code
   * @type integer
   */
  status?: (number | null) | undefined
  /**
   * @description Short title of the error
   * @type string
   */
  title?: (string | null) | undefined
  /**
   * @description URL of the error type. Can be used to lookup the error in a documentation
   * @type string
   */
  type?: (string | null) | undefined
}