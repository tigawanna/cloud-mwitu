import type {
  POSTCaddyHeaderParams,
  POSTCaddy201,
  POSTCaddy400,
  POSTCaddy500,
  POSTCaddyError,
  POSTCaddyMutationRequest,
  POSTCaddyMutationResponse,
} from "../../types/'caddyController/POSTCaddy.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { requestUpdateCaddyModelSchema } from '../requestUpdateCaddyModelSchema.ts'
import { updateCaddyResponseSchema } from '../updateCaddyResponseSchema.ts'
import { z } from 'zod'

export const POSTCaddyHeaderParamsSchema = z
  .object({
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<POSTCaddyHeaderParams>

/**
 * @description Created
 */
export const POSTCaddy201Schema = z.lazy(() => updateCaddyResponseSchema) as unknown as ToZod<POSTCaddy201>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const POSTCaddy400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<POSTCaddy400>

/**
 * @description Internal Server Error _(panics)_
 */
export const POSTCaddy500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<POSTCaddy500>

export const POSTCaddyErrorSchema = z.unknown() as unknown as ToZod<POSTCaddyError>

/**
 * @description Request body for controller.RequestUpdateCaddyModel
 */
export const POSTCaddyMutationRequestSchema = z.lazy(() => requestUpdateCaddyModelSchema) as unknown as ToZod<POSTCaddyMutationRequest>

export const POSTCaddyMutationResponseSchema = z.lazy(() => POSTCaddy201Schema) as unknown as ToZod<POSTCaddyMutationResponse>