import request from '@/utils/request'

export function listArticle (params) {
  return request({
    url: '/api/v1/articles',
    method: 'get',
    params
  })
}

export function getArticle (id) {
  return request({
    url: '/api/v1/article',
    method: 'get',
    params: { id }
  })
}

export function createArticle (data) {
  return request({
    url: '/api/v1/article',
    method: 'post',
    data
  })
}

export function updateArticle (data) {
  return request({
    url: '/api/v1/article',
    method: 'put',
    data
  })
}