import request from '@/utils/request'

export function listArticle (params) {
  return request({
    url: '/api/v1/articles',
    method: 'get',
    params
  })
}
