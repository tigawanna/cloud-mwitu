import type { HTTPError } from '../HTTPError.ts'
import type { SystemDService } from '../SystemDService.ts'

export type GETSystemdNamePathParams = {
  /**
   * @type string
   */
  name: string
}

export type GETSystemdNameQueryParams = {
  /**
   * @description look under /lib or /etc
   * @type string
   */
  libDir?: (string | null) | undefined
}

export type GETSystemdNameHeaderParams = {
  /**
   * @type string | undefined
   */
  Accept?: string | undefined
}

/**
 * @description OK
 */
export type GETSystemdName200 = SystemDService

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETSystemdName400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETSystemdName500 = HTTPError

export type GETSystemdNameError = unknown

export type GETSystemdNameQueryResponse = GETSystemdName200

export type GETsystemdnameQuery = {
  Response: GETSystemdName200
  PathParams: GETSystemdNamePathParams
  QueryParams: GETSystemdNameQueryParams
  HeaderParams: GETSystemdNameHeaderParams
  Errors: GETSystemdName400 | GETSystemdName500
}