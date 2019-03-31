import puppeteer from 'puppeteer';
// import devices from 'puppeteer/DeviceDescriptors';

class Browser{
  /**
   * 创建浏览器
   */
  public async Create(): Promise<any> {
    const options = {args: [
                      '--no-sandbox', 
                      '--disable-setuid-sandbox', 
                      '--process-per-tab', 
                      '--disable-images'
                    ],
                    headless: true}; // headless true 无头浏览器
    return await puppeteer.launch(options);
  }
}
export default new Browser();