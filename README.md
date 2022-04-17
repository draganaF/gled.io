# gled.io
Web aplikacija za kupovinu i rezervaciju bioskopskih karata. 

## Funkcionalnosti

1. Neregistrovani korisnik
- Može da se registruje, putem mail-a dobija aktivacioni link.
- Pregled i pretraga projekcija po datumu i vremenu projekcije, nazivu, žanru, glumacima, imenu režisera, sale, trajanju, godini proizvodnje, zemlji porijekla, ocjeni

2. Registrovani korisnik
- Pregled profila i izmjena osnovnih podataka
- Rezervacija karata 
- Kupovina karata (korisnik ima novčano stanje na profilu, koje može da dopunjava na kasi u bioskopu)
- Rezervacija i kupovina se obavljaju nakon što korisnik izabere sedišta preko prikaza sale za željenu projekciju
- Ocjenjivanje i ostavljanje recenzija na pregledane projekcije
- Prijavljivanje rezencija

3. Radnik
- Mijenja stanje računa korisnika nakon čega se korisnik obavještava putem email-a.
- Prodaja karata (uključuje i rezervisane karte)
- Dodaje projekcije i uklanja ih iz sistema

4. Admin
- Dodavanje zaposlenih
- Pregled i pretraga korisnika po imenu, prezimenu, email-u, roli u sistemu, broju kupljenih/rezervisanih karata, broju prodanih karata
- Pregled prijavljenih recenzija nakon čega admin može da blokira korisnika ukoliko uvidi nedolično ponašanje (putem email-a korisnik biva obavješten o blokiranju)

*Sistem će automatski pola sata prije početka projekcije ukloniti sve rezervacije koje nisu preuzete (kupljene) i korisniku se dodjelju negativan bod, nakon 3 negativna boda, korisnik biva blokiran.

## Arhitektura 
- User service - Go
- Projection service - Go
- Tickets Service - Python
- Recension service - Rust
- Email service - Go
- React klijentska aplikacija

*Podaci će biti čuvani u SQL bazi
