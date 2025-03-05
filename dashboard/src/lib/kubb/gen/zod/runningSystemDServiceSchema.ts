import type { RunningSystemDService } from '../types/RunningSystemDService.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description RunningSystemDService schema
 */
export const runningSystemDServiceSchema = z
  .object({
    activeState: z.string().optional(),
    loadState: z.string().optional(),
    name: z.string().optional(),
    subState: z.string().optional(),
    unit: z.string().optional(),
  })
  .describe('RunningSystemDService schema') as unknown as ToZod<RunningSystemDService>