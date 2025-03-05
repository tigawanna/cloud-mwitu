import type { HTTPError } from '../HTTPError.ts'
import type { RequestUpdateSystemDModel } from '../RequestUpdateSystemDModel.ts'
import type { SystemdServiceConfig } from '../SystemdServiceConfig.ts'

export type POSTSystemdHeaderParams = {
  /**
   * @type string | undefined
   */
  Accept?: string | undefined
}

/**
 * @description Created
 */
export type POSTSystemd201 = SystemdServiceConfig

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type POSTSystemd400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type POSTSystemd500 = HTTPError

export type POSTSystemdError = unknown

/**
 * @description Request body for controller.RequestUpdateSystemDModel
 */
export type POSTSystemdMutationRequest = RequestUpdateSystemDModel

export type POSTSystemdMutationResponse = POSTSystemd201

export type POSTsystemdMutation = {
  Response: POSTSystemd201
  Request: POSTSystemdMutationRequest
  HeaderParams: POSTSystemdHeaderParams
  Errors: POSTSystemd400 | POSTSystemd500
}