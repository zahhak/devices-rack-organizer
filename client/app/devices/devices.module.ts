import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular";
import { DeviceDetailComponent } from "./device-detail/device-detail.component";
import { DevicesListComponent } from "./device-list.component";
import { DevicesRoutingModule } from "./devices-routing.module";
import { DevicesService } from "~/shared/devices-service";

@NgModule({
    imports: [
        DevicesRoutingModule,
        NativeScriptCommonModule,
        NativeScriptFormsModule,
        NativeScriptUIListViewModule
    ],
    declarations: [
        DevicesListComponent,
        DeviceDetailComponent,
    ],
    providers: [
        DevicesService
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class DevicesModule { }
