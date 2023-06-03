import request from '@renderer/util/request'
import { AxiosResponse } from 'axios'
const userApi = {
  Login: '/user/login'
}
export function login(parameter): Promise<AxiosResponse> {
  return request({
    url: userApi.Login,
    method: 'post',
    data: parameter
  })
}
