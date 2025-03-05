import type { NoContent } from '../types/NoContent.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description NoContent schema
 */
export const noContentSchema = z
  .object({
    error: z.string().optional(),
    result: z.string().optional(),
  })
  .describe('NoContent schema') as unknown as ToZod<NoContent>