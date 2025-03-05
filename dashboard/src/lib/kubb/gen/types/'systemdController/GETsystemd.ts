import type { HTTPError } from '../HTTPError.ts'
import type { NoContent } from '../NoContent.ts'
import type { SystemDService } from '../SystemDService.ts'

export type GETSystemdQueryParams = {
  /**
   * @description Filter by name
   * @type string
   */
  name?: (string | null) | undefined
  /**
   * @description look under /lib or /etc
   * @type string
   */
  libDir?: (string | null) | undefined
}

export type GETSystemdHeaderParams = {
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
export type GETSystemd200 = SystemDService[]

/**
 * @description No Content
 */
export type GETSystemd204 = NoContent

/**
 * @description OK
 */
export type GETSystemd206 = unknown

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETSystemd400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETSystemd500 = HTTPError

export type GETSystemdError = unknown

export type GETSystemdQueryResponse = GETSystemd200 | GETSystemd204 | GETSystemd206

export type GETsystemdQuery = {
  Response: GETSystemd200 | GETSystemd204 | GETSystemd206
  QueryParams: GETSystemdQueryParams
  HeaderParams: GETSystemdHeaderParams
  Errors: GETSystemd400 | GETSystemd500
}