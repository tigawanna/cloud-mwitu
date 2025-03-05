import type {
  GETSystemdRunningQueryParams,
  GETSystemdRunningHeaderParams,
  GETSystemdRunning200,
  GETSystemdRunning400,
  GETSystemdRunning500,
  GETSystemdRunningError,
  GETSystemdRunningQueryResponse,
} from "../../types/'systemdController/GETSystemdRunning.ts"
import type { ToZod } from '@kubb/plugin-zod/utils'
import { HTTPErrorSchema } from '../HTTPErrorSchema.ts'
import { runningSystemDServiceSchema } from '../runningSystemDServiceSchema.ts'
import { z } from 'zod'

export const GETSystemdRunningQueryParamsSchema = z
  .object({
    name: z.string().describe('Filter by name').nullable().nullish(),
    libDir: z.string().describe('look under /lib or /etc').nullable().nullish(),
  })
  .optional() as unknown as ToZod<GETSystemdRunningQueryParams>

export const GETSystemdRunningHeaderParamsSchema = z
  .object({
    Accept: z.string().optional(),
  })
  .optional() as unknown as ToZod<GETSystemdRunningHeaderParams>

/**
 * @description OK
 */
export const GETSystemdRunning200Schema = z.array(z.lazy(() => runningSystemDServiceSchema)) as unknown as ToZod<GETSystemdRunning200>

/**
 * @description Bad Request _(validation or deserialization error)_
 */
export const GETSystemdRunning400Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemdRunning400>

/**
 * @description Internal Server Error _(panics)_
 */
export const GETSystemdRunning500Schema = z.lazy(() => HTTPErrorSchema) as unknown as ToZod<GETSystemdRunning500>

export const GETSystemdRunningErrorSchema = z.unknown() as unknown as ToZod<GETSystemdRunningError>

export const GETSystemdRunningQueryResponseSchema = z.lazy(() => GETSystemdRunning200Schema) as unknown as ToZod<GETSystemdRunningQueryResponse>