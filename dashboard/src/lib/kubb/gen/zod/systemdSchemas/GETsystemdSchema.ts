import type {
  GETSystemdQueryParams,
  GETSystemdHeaderParams,
  GETSystemd200,
  GETSystemd204,
  GETSystemd206,
  GETSystemd400,
  GETSystemd500,
  GETSystemdError,
  GETSystemdQueryResponse,
} from "../../types/'systemdController/GETSystemd.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { noContentSchema } from '../noContentSchema.ts'
import { systemDServiceSchema } from '../systemDServiceSchema.ts'
import { z } from 'zod'

export const GETSystemdQueryParamsSchema = z
  .object({
    name: z.string().describe('Filter by name').nullable().nullish(),
    libDir: z.string().describe('look under /lib or /etc').nullable().nullish(),
  })
  .optional() as unknown as ToZod<GETSystemdQueryParams>

export const GETSystemdHeaderParamsSchema = z
  .object({
    'X-Header': z.string().describe('header description').optional(),
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETSystemdHeaderParams>

/**
 * @description OK
 */
export const GETSystemd200Schema = z.array(z.lazy(() => systemDServiceSchema)) as unknown as ToZod<GETSystemd200>

/**
 * @description No Content
 */
export const GETSystemd204Schema = z.lazy(() => noContentSchema) as unknown as ToZod<GETSystemd204>

/**
 * @description OK
 */
export const GETSystemd206Schema = z.unknown() as unknown as ToZod<GETSystemd206>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETSystemd400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemd400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETSystemd500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemd500>

export const GETSystemdErrorSchema = z.unknown() as unknown as ToZod<GETSystemdError>

export const GETSystemdQueryResponseSchema = z.union([
  z.lazy(() => GETSystemd200Schema),
  z.lazy(() => GETSystemd204Schema),
  z.lazy(() => GETSystemd206Schema),
]) as unknown as ToZod<GETSystemdQueryResponse>