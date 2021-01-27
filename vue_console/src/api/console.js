import request from '@/utils/request'

export function getprojectinfo(data) {
    return request({
      url: '/lego/console/getprojectinfo',
      method: 'post',
      data
    })
}

export function gethostinfo(data) {
  return request({
    url: '/lego/console/gethostinfo',
    method: 'post',
    data
  })
}

export function getcpuinfo(data) {
  return request({
    url: '/lego/console/getcpuinfo',
    method: 'post',
    data
  })
}

export function getmemoryinfo(data) {
  return request({
    url: '/lego/console/getmemoryinfo',
    method: 'post',
    data
  })
}

export function gethostmonitordata(data) {
  return request({
    url: '/lego/console/gethostmonitordata',
    method: 'post',
    data
  })
}


export function getconsolecluster(data) {
  return request({
    url: '/lego/console/getconsolecluster',
    method: 'post',
    data
  })
}

export function setservicesetting(data) {
  return request({
    url: '/lego/console/setservicesetting',
    method: 'post',
    data
  })
}


export function setmodulesetting(data) {
  return request({
    url: '/lego/console/setmodulesetting',
    method: 'post',
    data
  })
}