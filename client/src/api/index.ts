import Axios from 'axios'

Axios.defaults.timeout = 180000

interface IConfig {
  baseURL: string,
  headers: Object
}

export default class Api {
  public static axios(_path: string, _data: any) {
    let fromData = new FormData()
    for (const key in _data) {
      if (_data.hasOwnProperty(key)) {
        const element = _data[key]
        fromData.append(key, element)
      }
    }

    /**
     * 配置参数
     */
    let config: IConfig = {
      baseURL: '127.0.0.1:3000/',
      headers: {
        'cms-channel': 0
      }
    }

    return Axios.request({
      method: 'POST',
      baseURL: config.baseURL,
      url: _path,
      data: fromData,
      headers: config.headers
    })
  }
}
