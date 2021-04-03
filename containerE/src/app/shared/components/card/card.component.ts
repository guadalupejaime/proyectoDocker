import { Component, Input, OnInit } from '@angular/core';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { characterGender, characterStatus } from '@pages/dashboard/characters/characters.models';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss']
})
export class CardComponent implements OnInit {
  @Input() character: Character;
  status = characterStatus;
  genders = characterGender;

  constructor() { }

  ngOnInit(): void {
  }

}
