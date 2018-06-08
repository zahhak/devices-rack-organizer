import { Injectable, NgZone } from "@angular/core";
import { Observable, throwError } from "rxjs";
import { catchError } from "rxjs/operators";
import { DeviceInfo } from "~/shared/device.model";
import { DeviceStatus } from "~/shared/deviceStatus.model";
import { HttpClient } from "~/shared/http-client";

@Injectable()
export class DevicesService {

    private _devices: DeviceInfo[];
    private serverUrl = "https://dp7o5mvps0.execute-api.eu-west-1.amazonaws.com/dev/device";

    private get devicesEndpoint(): string {
        return `${this.serverUrl}`;
    }

    private getDeviceActionEndpoint(deviceId): string {
        return `${this.serverUrl}/${deviceId}`;
    }

    constructor() { }

    public updateDeviceInfo(deviceId: string, user: string): Promise<DeviceStatus> {
        return HttpClient.call(this.getDeviceActionEndpoint(deviceId), "PUT", user)
            .then(res => {
                const response = res.content.toJSON();
                if (res.statusCode >= 400) {
                    throw new Error(response.message)
                }

                return response
            });
    }

    public getAllDevices(user: string): Promise<DeviceInfo[]> {
        return HttpClient.call(this.devicesEndpoint, "GET", user)
            .then(res => {
                this._devices = res.content.toJSON();
                return this._devices;
            });
    }

    public getDeviceInfo(user: string, deviceId: string): Promise<DeviceInfo> {
        const getDevicesPromise = (!this._devices || this._devices.length === 0) ?
            this.getAllDevices(user) : Promise.resolve(this._devices);

        return getDevicesPromise.then(res => {
            console.log("################### devices in getDeviceInfo", this._devices);
            const devicesFound = this._devices.filter(d => d.id !== deviceId);
            console.log("dev found", devicesFound);
            return devicesFound && devicesFound[0];
        });
    }
}
