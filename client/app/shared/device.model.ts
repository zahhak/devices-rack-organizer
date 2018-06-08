export class DeviceInfo {
    deviceId: string;
    userId: string;
    imageUrl: string;
    historyUrl: string;
    history?: any[];

    constructor(options: any) {
        this.deviceId = options.deviceId;
        this.userId = options.userId;
        this.imageUrl = options.imageUrl;
        this.historyUrl = options.historyUrl;
    }
}
