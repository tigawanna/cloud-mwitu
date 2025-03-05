import type { CaddyFileModel } from '../CaddyFileModel.ts'
import type { HTTPError } from '../HTTPError.ts'
import type { NoContent } from '../NoContent.ts'

export type GETCaddyNamePathParams = {
  /**
   * @type string
   */
  name: string
}

export type GETCaddyNameHeaderParams = {
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
export type GETCaddyName200 = CaddyFileModel

/**
 * @description No Content
 */
export type GETCaddyName204 = NoContent

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETCaddyName400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETCaddyName500 = HTTPError

export type GETCaddyNameError = unknown

export type GETCaddyNameQueryResponse = GETCaddyName200 | GETCaddyName204

export type GETcaddynameQuery = {
  Response: GETCaddyName200 | GETCaddyName204
  PathParams: GETCaddyNamePathParams
  HeaderParams: GETCaddyNameHeaderParams
  Errors: GETCaddyName400 | GETCaddyName500
}