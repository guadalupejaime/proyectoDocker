import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ModalTemplateComponent } from '@shared/components/modal-template/modal-template.component';
import { MyMaterialModule } from '@shared/my-material/my-material.module';



@NgModule({
  declarations: [
    ModalTemplateComponent
  ],
  exports: [
    ModalTemplateComponent
  ],
  imports: [
    CommonModule,
    MyMaterialModule
  ]
})
export class ModalTemplateModule { }
