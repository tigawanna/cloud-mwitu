import type { UpdateCaddyResponse } from '../types/UpdateCaddyResponse.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description UpdateCaddyResponse schema
 */
export const updateCaddyResponseSchema = z
  .object({
    content: z.string().optional(),
    contentArray: z.array(z.string()).optional(),
    contentArrayBefore: z.array(z.string()).optional(),
    updatedBlock: z.string().optional(),
  })
  .describe('UpdateCaddyResponse schema') as unknown as ToZod<UpdateCaddyResponse>