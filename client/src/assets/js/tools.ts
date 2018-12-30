import 'isomorphic-fetch'

export default class Tools {
  private static extractIp(rawText: string) {
    return rawText.replace(/\s+/g, '')
  }

  public static async getIp(): Promise<string> {
    const request = await fetch('https://ipv4.icanhazip.com/')
    const text = await request.text()
    const ip = this.extractIp(text)
    return ip
  }
}
