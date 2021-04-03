import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LocationsComponent } from '@pages/dashboard/locations/locations.component';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';
import { ModalTemplateModule } from '@shared/components/modal-template/modal-template.module';
import { MyMaterialModule } from '@shared/my-material/my-material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { InfoLocationComponent } from '@pages/dashboard/locations/modals/info-location/info-location.component';
import { AddLocationComponent } from '@pages/dashboard/locations/modals/add-location/add-location.component';



@NgModule({
  declarations: [
    LocationsComponent,
    InfoLocationComponent,
    AddLocationComponent
  ],
  entryComponents: [
    InfoLocationComponent,
    AddLocationComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    PageTemplateModule,
    ModalTemplateModule,
    MyMaterialModule
  ]
})
export class LocationsModule { }
