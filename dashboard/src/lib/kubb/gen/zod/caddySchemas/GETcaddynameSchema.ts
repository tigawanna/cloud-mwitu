import type {
  GETCaddyNamePathParams,
  GETCaddyNameHeaderParams,
  GETCaddyName200,
  GETCaddyName204,
  GETCaddyName400,
  GETCaddyName500,
  GETCaddyNameError,
  GETCaddyNameQueryResponse,
} from "../../types/'caddyController/GETCaddyName.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { caddyFileModelSchema } from '../caddyFileModelSchema.ts'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { noContentSchema } from '../noContentSchema.ts'
import { z } from 'zod'

export const GETCaddyNamePathParamsSchema = z.object({
  name: z.string(),
}) as unknown as ToZod<GETCaddyNamePathParams>

export const GETCaddyNameHeaderParamsSchema = z
  .object({
    'X-Header': z.string().describe('header description').optional(),
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETCaddyNameHeaderParams>

/**
 * @description OK
 */
export const GETCaddyName200Schema = z.lazy(() => caddyFileModelSchema) as unknown as ToZod<GETCaddyName200>

/**
 * @description No Content
 */
export const GETCaddyName204Schema = z.lazy(() => noContentSchema) as unknown as ToZod<GETCaddyName204>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETCaddyName400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddyName400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETCaddyName500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddyName500>

export const GETCaddyNameErrorSchema = z.unknown() as unknown as ToZod<GETCaddyNameError>

export const GETCaddyNameQueryResponseSchema = z.union([
  z.lazy(() => GETCaddyName200Schema),
  z.lazy(() => GETCaddyName204Schema),
]) as unknown as ToZod<GETCaddyNameQueryResponse>