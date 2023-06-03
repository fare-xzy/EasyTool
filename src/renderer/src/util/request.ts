import axios from 'axios'

export interface ReturnResult<T> {
  code: number
  message: string
  data: T
}
const request = axios.create({
  baseURL: import.meta.env.RENDERER_VITE_BASE_URL,
  timeout: 20000
})

export default request
