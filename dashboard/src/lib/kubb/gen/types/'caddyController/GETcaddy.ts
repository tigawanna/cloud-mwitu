import type { CaddyFileModel } from '../CaddyFileModel.ts'
import type { HTTPError } from '../HTTPError.ts'

export type GETCaddyQueryParams = {
  /**
   * @description Filter by name
   * @type string
   */
  name?: (string | null) | undefined
}

export type GETCaddyHeaderParams = {
  /**
   * @type string | undefined
   */
  Accept?: string | undefined
}

/**
 * @description OK
 */
export type GETCaddy200 = CaddyFileModel[]

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETCaddy400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETCaddy500 = HTTPError

export type GETCaddyError = unknown

export type GETCaddyQueryResponse = GETCaddy200

export type GETcaddyQuery = {
  Response: GETCaddy200
  QueryParams: GETCaddyQueryParams
  HeaderParams: GETCaddyHeaderParams
  Errors: GETCaddy400 | GETCaddy500
}