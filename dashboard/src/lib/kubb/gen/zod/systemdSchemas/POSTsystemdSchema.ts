import type {
  POSTSystemdHeaderParams,
  POSTSystemd201,
  POSTSystemd400,
  POSTSystemd500,
  POSTSystemdError,
  POSTSystemdMutationRequest,
  POSTSystemdMutationResponse,
} from "../../types/'systemdController/POSTSystemd.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { requestUpdateSystemDModelSchema } from '../requestUpdateSystemDModelSchema.ts'
import { systemdServiceConfigSchema } from '../systemdServiceConfigSchema.ts'
import { z } from 'zod'

export const POSTSystemdHeaderParamsSchema = z
  .object({
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<POSTSystemdHeaderParams>

/**
 * @description Created
 */
export const POSTSystemd201Schema = z.lazy(() => systemdServiceConfigSchema) as unknown as ToZod<POSTSystemd201>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const POSTSystemd400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<POSTSystemd400>

/**
 * @description Internal Server Error _(panics)_
 */
export const POSTSystemd500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<POSTSystemd500>

export const POSTSystemdErrorSchema = z.unknown() as unknown as ToZod<POSTSystemdError>

/**
 * @description Request body for controller.RequestUpdateSystemDModel
 */
export const POSTSystemdMutationRequestSchema = z.lazy(() => requestUpdateSystemDModelSchema) as unknown as ToZod<POSTSystemdMutationRequest>

export const POSTSystemdMutationResponseSchema = z.lazy(() => POSTSystemd201Schema) as unknown as ToZod<POSTSystemdMutationResponse>