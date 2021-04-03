export interface Location {
    created: string;
    dimension: string;
    id: number;
    name: string;
    residents: string[];
    type: string;
    url: string;
}

export interface Info {
    count: number;
    next: string;
    page: number;
    prev: string;
}

export interface Locations {
    info: Info;
    results: Location[];
    total_found: number;
    total_returned: number;
}
