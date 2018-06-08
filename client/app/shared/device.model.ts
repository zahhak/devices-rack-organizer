export class DeviceInfo {
    deviceId: string;
    userId: string;
    imageUrl: string;
    historyUrl: string;

    constructor(options: any) {
        this.deviceId = options.deviceId;
        this.userId = options.userId;
        this.imageUrl = options.imageUrl;
        this.historyUrl = options.historyUrl;
    }
}
