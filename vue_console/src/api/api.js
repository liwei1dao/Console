import request from '@/utils/request'

export function sendemailcaptcha(data) {
    return request({
      url: '/lego/api/sendemailcaptcha',
      method: 'post',
      data
    })
}

