import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CharactersComponent } from '@pages/dashboard/characters/characters.component';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';
import { MyMaterialModule } from '@shared/my-material/my-material.module';
import { AddCharacterComponent } from './modals/add-character/add-character.component';
import { InfoCharacterComponent } from '@pages/dashboard/characters/modals/info-character/info-character.component';
import { ModalTemplateModule } from '@shared/components/modal-template/modal-template.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [
    CharactersComponent,
    AddCharacterComponent,
    InfoCharacterComponent
  ],
  entryComponents: [
    AddCharacterComponent,
    InfoCharacterComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    PageTemplateModule,
    MyMaterialModule,
    ModalTemplateModule
  ],
})
export class CharactersModule { }
