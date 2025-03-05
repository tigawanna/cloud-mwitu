import type { NoContent } from '../types/NoContent.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description NoContent schema
 */
export const noContentSchema = z.unknown() as unknown as ToZod<NoContent>