import type { HTTPError } from '../HTTPError.ts'
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
   * @type string | undefined
   */
  Accept?: string | undefined
}

/**
 * @description OK
 */
export type GETSystemdRunning200 = RunningSystemDService[]

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type GETSystemdRunning400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type GETSystemdRunning500 = HTTPError

export type GETSystemdRunningError = unknown

export type GETSystemdRunningQueryResponse = GETSystemdRunning200

export type GETsystemdrunningQuery = {
  Response: GETSystemdRunning200
  QueryParams: GETSystemdRunningQueryParams
  HeaderParams: GETSystemdRunningHeaderParams
  Errors: GETSystemdRunning400 | GETSystemdRunning500
}