import type {
  GETSystemdNamePathParams,
  GETSystemdNameQueryParams,
  GETSystemdNameHeaderParams,
  GETSystemdName200,
  GETSystemdName400,
  GETSystemdName500,
  GETSystemdNameError,
  GETSystemdNameQueryResponse,
} from "../../types/'systemdController/GETSystemdName.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { systemDServiceSchema } from '../systemDServiceSchema.ts'
import { z } from 'zod'

export const GETSystemdNamePathParamsSchema = z.object({
  name: z.string(),
}) as unknown as ToZod<GETSystemdNamePathParams>

export const GETSystemdNameQueryParamsSchema = z
  .object({
    libDir: z.string().describe('look under /lib or /etc').nullable().nullish(),
  })
  .optional() as unknown as ToZod<GETSystemdNameQueryParams>

export const GETSystemdNameHeaderParamsSchema = z
  .object({
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETSystemdNameHeaderParams>

/**
 * @description OK
 */
export const GETSystemdName200Schema = z.lazy(() => systemDServiceSchema) as unknown as ToZod<GETSystemdName200>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETSystemdName400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemdName400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETSystemdName500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemdName500>

export const GETSystemdNameErrorSchema = z.unknown() as unknown as ToZod<GETSystemdNameError>

export const GETSystemdNameQueryResponseSchema = z.lazy(() => GETSystemdName200Schema) as unknown as ToZod<GETSystemdNameQueryResponse>