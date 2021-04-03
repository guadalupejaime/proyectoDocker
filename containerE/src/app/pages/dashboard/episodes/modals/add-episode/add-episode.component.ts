import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { Months } from '@pages/dashboard/characters/characters.models';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-add-episode',
  templateUrl: './add-episode.component.html',
  styleUrls: ['./add-episode.component.scss']
})
export class AddEpisodeComponent implements OnInit {
  @ViewChild('characterInput') characterInput: ElementRef<HTMLInputElement>;
  formEpisode: FormGroup;
  selectedCharacters: Character[] = [];
  filteredCharacters: Character[] = [];
  ctrlCharacters = new FormControl([], Validators.required);
  characters: Character[] = [];
  today = new Date();

  constructor(
    public sMatDialogRef: MatDialogRef<AddEpisodeComponent>,
    private sHttpService: HttpService
  ) {
    this.formEpisode = new FormGroup({
      name: new FormControl(null, Validators.required),
      air_date: new FormControl(this.today, Validators.required),
      episode: new FormControl(null, Validators.required),
    });
  }

  ngOnInit(): void {
    this.getCharacters();
    this.ctrlCharacters.valueChanges.subscribe(
      (character: string) => {
        const temp = typeof character;
        if (temp === 'string') {
          this.filterCharacters(character.toLowerCase());
        }
      });
  }

  // get characters list
  private getCharacters(): void {
    this.sHttpService.getCharacters({ limit: 500, offset: 0 }).subscribe(
      (characters) => {
        characters.results.forEach(character => {
          const newCharacter: Character = { ...character, checked: false };
          this.characters.push(newCharacter);
        });
      },
      (err) => {

      }
    );
  }

  private filterCharacters(myFilter?: string): void {
    if (myFilter) {
      this.filteredCharacters = this.characters.filter((character) => {
        return character.name.toLowerCase().indexOf(myFilter) === 0;
      });
    } else {
      if (this.characters.length < 15) {
        this.filteredCharacters = this.characters;
      }
    }
  }

  displayFnCharacter(character: Character): string {
    return character && character.name ? character.name : '';
  }

  filterCharacterSelected(character: Character): void {
    const index = this.filteredCharacters.indexOf(character);
    if (index >= 0) {
      this.filteredCharacters[index].checked = true;
    }
    if (this.selectedCharacters.indexOf(character) < 0) {
      this.selectedCharacters.push(character);
    }
    this.ctrlCharacters.setValue(this.selectedCharacters);
  }

  removeEpisode(character: Character): void {
    const index = this.selectedCharacters.indexOf(character);

    if (index >= 0) {
      this.selectedCharacters.splice(index, 1);
      this.filteredCharacters[index].checked = false;
    }
  }

  changeCharacters(character: Character): void {
    const index = this.filteredCharacters.indexOf(character);
    this.characterInput.nativeElement.value = '';
    if (index >= 0) {
      if (this.filteredCharacters[index].checked) {
        this.removeEpisode(character);
        return;
      }
      this.filterCharacterSelected(character);
    }
  }

  isValidForm(): boolean {
    const validCharacters = this.ctrlCharacters.value.length > 0;
    this.formEpisode.markAllAsTouched();
    this.ctrlCharacters.markAsTouched();
    return this.formEpisode.valid && validCharacters;
  }

  formatDate(date: Date): string {
    let newDate = '';
    newDate += Months[date.getMonth()];
    newDate += ' ' + date.getDate();
    newDate += ', ' + date.getFullYear();

    return newDate;
  }

  postEpisode(): void {
    if (this.isValidForm()) {
      const charactesSelected = [];
      for (const character of this.ctrlCharacters.value) {
        charactesSelected.push(character.id + '');
      }
      const newEpisode = {
        ...this.formEpisode.value,
        characters: charactesSelected
      };
      const date = this.formEpisode.get('air_date').value;
      newEpisode.air_date = this.formatDate(date);

      console.log('new-episode', newEpisode);

      this.sHttpService.postEpisode(newEpisode).subscribe(
        (resp) => {
          this.sMatDialogRef.close(true);
        },
        (err) => {
          console.log(err);
          this.sMatDialogRef.close('error');
        }
      );
    }
  }

}
