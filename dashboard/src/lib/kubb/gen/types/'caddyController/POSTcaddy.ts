import type { HTTPError } from '../HTTPError.ts'
import type { NoContent } from '../NoContent.ts'
import type { RequestUpdateCaddyModel } from '../RequestUpdateCaddyModel.ts'
import type { UpdateCaddyResponse } from '../UpdateCaddyResponse.ts'

export type POSTCaddyHeaderParams = {
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
 * @description Created
 */
export type POSTCaddy201 = UpdateCaddyResponse

/**
 * @description No Content
 */
export type POSTCaddy204 = NoContent

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export type POSTCaddy400 = HTTPError

/**
 * @description Internal Server Error _(panics)_
 */
export type POSTCaddy500 = HTTPError

export type POSTCaddyError = unknown

/**
 * @description Request body for controller.RequestUpdateCaddyModel
 */
export type POSTCaddyMutationRequest = RequestUpdateCaddyModel

export type POSTCaddyMutationResponse = POSTCaddy201 | POSTCaddy204

export type POSTcaddyMutation = {
  Response: POSTCaddy201 | POSTCaddy204
  Request: POSTCaddyMutationRequest
  HeaderParams: POSTCaddyHeaderParams
  Errors: POSTCaddy400 | POSTCaddy500
}