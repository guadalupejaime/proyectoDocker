import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-add-location',
  templateUrl: './add-location.component.html',
  styleUrls: ['./add-location.component.scss']
})
export class AddLocationComponent implements OnInit {
  @ViewChild('locationInput') locationInput: ElementRef<HTMLInputElement>;
  formLocation: FormGroup;
  selectedCharacters: Character[] = [];
  filteredCharacters: Character[] = [];
  ctrlCharacters = new FormControl([], Validators.required);
  locations: Character[] = [];

  constructor(
    public sMatDialogRef: MatDialogRef<AddLocationComponent>,
    private sHttpService: HttpService
  ) {
    this.formLocation = new FormGroup({
      name: new FormControl(null, Validators.required),
      type: new FormControl(null, Validators.required),
      dimension: new FormControl(null, Validators.required),
    });
  }

  ngOnInit(): void {
    this.getCharacters();
    this.ctrlCharacters.valueChanges.subscribe(
      (location: string) => {
        const temp = typeof location;
        if (temp === 'string') {
          this.filterCharacters(location.toLowerCase());
        }
      });
  }

  // get locations list
  private getCharacters(): void {
    this.sHttpService.getCharacters({ limit: 500, offset: 0 }).subscribe(
      (locations) => {
        locations.results.forEach(location => {
          const newCharacter: Character = { ...location, checked: false };
          this.locations.push(newCharacter);
        });
      },
      (err) => {

      }
    );
  }

  private filterCharacters(myFilter?: string): void {
    if (myFilter) {
      this.filteredCharacters = this.locations.filter((location) => {
        return location.name.toLowerCase().indexOf(myFilter) === 0;
      });
    } else {
      if (this.locations.length < 15) {
        this.filteredCharacters = this.locations;
      }
    }
  }

  displayFnCharacter(location: Character): string {
    return location && location.name ? location.name : '';
  }

  filterCharacterSelected(location: Character): void {
    const index = this.filteredCharacters.indexOf(location);
    if (index >= 0) {
      this.filteredCharacters[index].checked = true;
    }
    if (this.selectedCharacters.indexOf(location) < 0) {
      this.selectedCharacters.push(location);
    }
    this.ctrlCharacters.setValue(this.selectedCharacters);
  }

  removeLocation(location: Character): void {
    const index = this.selectedCharacters.indexOf(location);

    if (index >= 0) {
      this.selectedCharacters.splice(index, 1);
      this.filteredCharacters[index].checked = false;
    }
  }

  changeCharacters(location: Character): void {
    const index = this.filteredCharacters.indexOf(location);
    this.locationInput.nativeElement.value = '';
    if (index >= 0) {
      if (this.filteredCharacters[index].checked) {
        this.removeLocation(location);
        return;
      }
      this.filterCharacterSelected(location);
    }
  }

  isValidForm(): boolean {
    const validCharacters = this.ctrlCharacters.value.length > 0;
    this.formLocation.markAllAsTouched();
    this.ctrlCharacters.markAsTouched();
    return this.formLocation.valid && validCharacters;
  }

  postLocation(): void {
    if (this.isValidForm()) {
      const charactesSelected = [];
      for (const location of this.ctrlCharacters.value) {
        charactesSelected.push(location.id + '');
      }
      const newLocation = {
        ...this.formLocation.value,
        url: 'https://rickandmortyapi.com/api/location',
        residents: charactesSelected
      };
      this.sHttpService.postLocation(newLocation).subscribe(
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
