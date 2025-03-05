import type { CaddyFileModel } from '../CaddyFileModel.ts'
import type { HTTPError } from '../HTTPError.ts'
import type { NoContent } from '../NoContent.ts'

export type GETCaddyQueryParams = {
  /**
   * @description Filter by name
   * @type string
   */
  name?: (string | null) | undefined
}

export type GETCaddyHeaderParams = {
  /**
   * @description header description
   * @type string | undefined
   */
  'X-Header'?: string | undefined
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
 * @description No Content
 */
export type GETCaddy204 = NoContent

/**
 * @description OK
 */
export type GETCaddy206 = unknown

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETCaddy400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETCaddy500 = HTTPError

export type GETCaddyError = unknown

export type GETCaddyQueryResponse = GETCaddy200 | GETCaddy204 | GETCaddy206

export type GETcaddyQuery = {
  Response: GETCaddy200 | GETCaddy204 | GETCaddy206
  QueryParams: GETCaddyQueryParams
  HeaderParams: GETCaddyHeaderParams
  Errors: GETCaddy400 | GETCaddy500
}