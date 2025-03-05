import type { HTTPError } from '../HTTPError.ts'
import type { NoContent } from '../NoContent.ts'
import type { RunningSystemDService } from '../RunningSystemDService.ts'

export type GETSystemdRunningQueryParams = {
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

export type GETSystemdRunningHeaderParams = {
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
export type GETSystemdRunning200 = RunningSystemDService[]

/**
 * @description No Content
 */
export type GETSystemdRunning204 = NoContent

/**
 * @description OK
 */
export type GETSystemdRunning206 = unknown

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETSystemdRunning400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETSystemdRunning500 = HTTPError

export type GETSystemdRunningError = unknown

export type GETSystemdRunningQueryResponse = GETSystemdRunning200 | GETSystemdRunning204 | GETSystemdRunning206

export type GETsystemdrunningQuery = {
  Response: GETSystemdRunning200 | GETSystemdRunning204 | GETSystemdRunning206
  QueryParams: GETSystemdRunningQueryParams
  HeaderParams: GETSystemdRunningHeaderParams
  Errors: GETSystemdRunning400 | GETSystemdRunning500
}