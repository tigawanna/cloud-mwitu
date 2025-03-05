import type {
  GETCaddyQueryParams,
  GETCaddyHeaderParams,
  GETCaddy200,
  GETCaddy204,
  GETCaddy206,
  GETCaddy400,
  GETCaddy500,
  GETCaddyError,
  GETCaddyQueryResponse,
} from "../../types/'caddyController/GETCaddy.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { caddyFileModelSchema } from '../caddyFileModelSchema.ts'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { noContentSchema } from '../noContentSchema.ts'
import { z } from 'zod'

export const GETCaddyQueryParamsSchema = z
  .object({
    name: z.string().describe('Filter by name').nullable().nullish(),
  })
  .optional() as unknown as ToZod<GETCaddyQueryParams>

export const GETCaddyHeaderParamsSchema = z
  .object({
    'X-Header': z.string().describe('header description').optional(),
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETCaddyHeaderParams>

/**
 * @description OK
 */
export const GETCaddy200Schema = z.array(z.lazy(() => caddyFileModelSchema)) as unknown as ToZod<GETCaddy200>

/**
 * @description No Content
 */
export const GETCaddy204Schema = z.lazy(() => noContentSchema) as unknown as ToZod<GETCaddy204>

/**
 * @description OK
 */
export const GETCaddy206Schema = z.unknown() as unknown as ToZod<GETCaddy206>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETCaddy400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddy400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETCaddy500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETCaddy500>

export const GETCaddyErrorSchema = z.unknown() as unknown as ToZod<GETCaddyError>

export const GETCaddyQueryResponseSchema = z.union([
  z.lazy(() => GETCaddy200Schema),
  z.lazy(() => GETCaddy204Schema),
  z.lazy(() => GETCaddy206Schema),
]) as unknown as ToZod<GETCaddyQueryResponse>