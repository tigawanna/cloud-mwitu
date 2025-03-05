export const operations = {
  'GET_/caddy/': {
    path: '/caddy/',
    method: 'get',
  },
  'POST_/caddy/': {
    path: '/caddy/',
    method: 'post',
  },
  'GET_/caddy/:name': {
    path: '/caddy/:name',
    method: 'get',
  },
  'GET_/systemd/': {
    path: '/systemd/',
    method: 'get',
  },
  'POST_/systemd/': {
    path: '/systemd/',
    method: 'post',
  },
  'GET_/systemd/running': {
    path: '/systemd/running',
    method: 'get',
  },
  'GET_/systemd/:name': {
    path: '/systemd/:name',
    method: 'get',
  },
} as const