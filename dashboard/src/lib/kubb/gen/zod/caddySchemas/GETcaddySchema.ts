import type {
  GETCaddyQueryParams,
  GETCaddyHeaderParams,
  GETCaddy200,
  GETCaddy400,
  GETCaddy500,
  GETCaddyError,
  GETCaddyQueryResponse,
} from "../../types/'caddyController/GETCaddy.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { caddyFileModelSchema } from '../caddyFileModelSchema.ts'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { z } from 'zod'

export const GETCaddyQueryParamsSchema = z
  .object({
    name: z.string().describe('Filter by name').nullable().nullish(),
  })
  .optional() as unknown as ToZod<GETCaddyQueryParams>

export const GETCaddyHeaderParamsSchema = z
  .object({
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETCaddyHeaderParams>

/**
 * @description OK
 */
export const GETCaddy200Schema = z.array(z.lazy(() => caddyFileModelSchema)) as unknown as ToZod<GETCaddy200>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETCaddy400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddy400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETCaddy500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddy500>

export const GETCaddyErrorSchema = z.unknown() as unknown as ToZod<GETCaddyError>

export const GETCaddyQueryResponseSchema = z.lazy(() => GETCaddy200Schema) as unknown as ToZod<GETCaddyQueryResponse>