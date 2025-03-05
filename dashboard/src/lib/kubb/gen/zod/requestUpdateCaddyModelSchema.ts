import type { RequestUpdateCaddyModel } from '../types/RequestUpdateCaddyModel.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description RequestUpdateCaddyModel schema
 */
export const requestUpdateCaddyModelSchema = z
  .object({
    content: z.string().optional(),
    name: z.string().optional(),
  })
  .describe('RequestUpdateCaddyModel schema') as unknown as ToZod<RequestUpdateCaddyModel>