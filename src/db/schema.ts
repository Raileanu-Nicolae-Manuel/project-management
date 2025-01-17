import { pgTable, varchar } from "drizzle-orm/pg-core";
import { defaultTable } from "./interface/types";

export default pgTable('users', {
  name: varchar({ length: 255 }),
  email: varchar({ length: 255 }).notNull().unique(),
  ...defaultTable,
});
