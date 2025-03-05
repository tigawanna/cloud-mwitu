import type { UnknownInterface } from '../types/UnknownInterface.ts'
import type { ToZod } from '@kubb/plugin-zod/utils'
import { z } from 'zod'

/**
 * @description unknown-interface schema
 */
export const unknownInterfaceSchema = z.unknown() as unknown as ToZod<UnknownInterface>