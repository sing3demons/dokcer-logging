import express from 'express'
import { createLogger, format, transports } from 'winston'

const app = express()
const port = 3000

const logger = createLogger({
  level: 'debug',
  format: format.simple(),
  // You can also comment out the line above and uncomment the line below for JSON format
  format: format.json(),
  transports: [new transports.Console()],
})

app.get('/', (req, res) => {
  logger.info('Hello World!')
  res.json({ message: 'Hello World!' })
})

app.listen(port, () => logger.info(`Express server is listening at http://localhost:${port} 🚀`))
