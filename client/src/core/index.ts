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
  public static reply (content: string) : string {
    return ''
  }

  /**
   * 学习中枢
   * @param {String} content
   */
  public static teach (content: string) : Object {
    return {}
  }

  /**
   * 记忆消除中枢
   * @param {Object[]} list
   */
  public static forget (list: Object[]) : Boolean {
    return false
  }
}
