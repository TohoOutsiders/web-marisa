import Api from '../api/v1'

interface ISpeakConfig {
  name: string
  content: string
}

export default class Core {
  /**
   * 魔理沙说话格式,以及处理You的说话格式
   * @param {String} name
   * @param {String} content
   */
  public static speak (name: string, content: string) : Object {
    let obj: ISpeakConfig = {
      name: name,
      content: content
    }
    return obj
  }

  /**
   * 回复逻辑判断中枢
   * @param {String} content
   */
  public static async reply (content: string) : Promise<any> {
    let config = {
      keyword: content
    }
    try {
      let res = await Api.fecthMemory(config)
      return res.data.data.answer
    } catch (err) {
      console.log(`回复失败 ... ${err}`)
    }
  }

  /**
   * 学习中枢
   * @param {String} content
   */
  public static async teach (content: string) : Promise<any> {
    let str = content.split('`')
    let config = {
      ip: '127.0.0.1',
      keyword: str[0],
      answer: str[1]
    }
    try {
      let res = await Api.AddMemory(config)
      if (res.data.data.code === 200) {
        return true
      }
    } catch (err) {
      console.log(`无法学习 ... ${err}`)
      return false
    }
  }

  /**
   * 记忆消除中枢
   * @param {Object[]} list
   */
  public static forget (list: Object[]) : Boolean {
    return false
  }
}
