import request from '@/utils/request'

export function listTag (params) {
  return request({
    url: '/api/v1/tags',
    method: 'get',
    params
  })
}

export function createTag (data) {
  return request({
    url: '/api/v1/tag',
    method: 'post',
    data
  })
}

export function updateTag (data) {
  return request({
    url: '/api/v1/tag',
    method: 'put',
    data
  })
}

export function getTag (id) {
  return request({
    url: '/api/v1/tag',
    method: 'get',
    params: { id }
  })
}

export function deleteTag (id) {
  return request({
    url: '/api/v1/tag',
    method: 'delete',
    params: { id }
  })
}