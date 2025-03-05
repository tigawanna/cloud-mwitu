import type { HTTPError } from '../types/HTTPError.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description HTTPError schema
 */
export const HTTPErrorSchema = z
  .object({
    detail: z.string().describe('Human readable error message').nullable().nullish(),
    errors: z
      .array(
        z.object({
          more: z.object({}).catchall(z.unknown()).optional(),
          name: z.string().optional(),
          reason: z.string().optional(),
        }),
      )
      .nullable()
      .nullish(),
    instance: z.string().nullable().nullish(),
    status: z.number().int().describe('HTTP status code').nullable().nullish(),
    title: z.string().describe('Short title of the error').nullable().nullish(),
    type: z.string().describe('URL of the error type. Can be used to lookup the error in a documentation').nullable().nullish(),
  })
  .describe('HTTPError schema') as unknown as ToZod<HTTPError>