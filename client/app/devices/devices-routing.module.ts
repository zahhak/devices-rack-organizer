import { NgModule } from "@angular/core";
import { Routes } from "@angular/router";
import { NativeScriptRouterModule } from "nativescript-angular/router";

import { DeviceDetailComponent } from "./device-detail/device-detail.component";
import { DevicesListComponent } from "./device-list.component";

const routes: Routes = [
    { path: "", component: DevicesListComponent },
    { path: "device-detail", component: DeviceDetailComponent },
];

@NgModule({
    imports: [NativeScriptRouterModule.forChild(routes)],
    exports: [NativeScriptRouterModule]
})
export class DevicesRoutingModule { }
