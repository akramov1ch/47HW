## Uy vazifa
### Protokol buferlari va `PostgreSQL` bilan `RESTful API` yaratish
### `API` uchta endpoint qo'llab-quvvatlaydi:
1. User `create` (`JSON`-ni qabul qiladi, protobufga marshal qiladi va ma'lumotlar bazasiga saqlaydi).
2. `Get` users da [protojson](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) `encoding` qilingan holatda qaytaring.
3. Get a user by ID.

### Proto file mazmuni:
1. `id` <br>
    Turi: int32 <br>
    Talab: Bu har bir shaxs uchun noyob identifikator. U ma'lumotlar bazasi tomonidan avtomatik ravishda yaratiladi va yaratilish paytida mijoz tomonidan taqdim etilmasligi kerak. <br>
    Misol: 1
2. `name` <br>
    Turi: string <br>
    Talab: Shaxsning to'liq ismi. Ushbu maydon majburiy va bo'sh bo'lmagan qator bo'lishi kerak. <br>
    Misol: "Falonchi Pistonchiev"
3. `age` <br>
    Turi: int32 <br>
    Talab: Shaxsning yoshi. Ushbu maydon majburiy va manfiy bo'lmagan butun son bo'lishi kerak. <br>
    Misol: 30
4. `email` <br>
    Turi: string <br>
    Talab: Shaxsning elektron pochta manzili. Ushbu maydon majburiy va haqiqiy elektron pochta formatida bo'lishi kerak. <br>
    Misol: "falonchi@example.com"
5. `address` <br>
    Turi: Address (ichki xabar)  <br>
    Talab: Ushbu ichki xabar shaxsning manzil ma'lumotlarini o'z ichiga oladi. Address xabari ichidagi barcha maydonlar majburiydir. <br>
    street: Ko'cha manzili (masalan, "123 Main St"). <br>
    city: Shahar (masalan, "Toshkent"). <br>
    zipcode: Pochta indeksi (masalan, "12345").
6. `phone_numbers` <br>
    Turi: repeated `PhoneNumber` (ichki xabar) <br>
    Talab: Shaxs bilan bog'liq telefon raqamlarining ro'yxati. Har bir telefon raqami ichki `PhoneNumber` xabari bilan taqdim etilishi kerak. Ushbu maydon ixtiyoriy, lekin agar taqdim etilgan bo'lsa, barcha ichki maydonlar haqiqiy bo'lishi kerak. <br>
    number: Telefon raqami (masalan, "(90)-123-45-67"). <br>
    type: Telefon raqami turi, masalan, "mobile", "home", yoki "work".
7. `occupation` <br>
    Turi: string <br>
    Talab: Shaxsning kasbi. Ushbu maydon ixtiyoriy, lekin taqdim etilgan bo'lsa, haqiqiy qator bo'lishi kerak. <br>
    Misol: "Software Engineer"
8. `company` <br>
    Turi: string <br>
    Talab: Shaxs bilan bog'liq kompaniya. Ushbu maydon ixtiyoriy, lekin taqdim etilgan bo'lsa, haqiqiy qator bo'lishi kerak. <br>
    Misol: "Tech Inc."
9. `is_active` <br>
    Turi: bool  <br>
    Talab: Shaxsning hozirgi faol ekanligini ko'rsatadigan mantiqiy bayroq. Ushbu maydon ixtiyoriy, lekin taqdim etilgan bo'lsa, true yoki false bo'lishi kerak. <br>
    Misol: true

### JSON Ko'rinishidagi Misol
```json
{
  "name": "Falonchi Palonchiev",
  "age": 30,
  "email": "falonchiev@example.com",
  "address": {
    "street": "123 Main St",
    "city": "Tashkent",
    "zipcode": "12345"
  },
  "phone_numbers": [
    {"number": "(90)-123-45-67", "type": "mobile"},
    {"number": "(71)-123-45-67", "type": "home"}
  ],
  "occupation": "Software Engineer",
  "company": "Tech Inc.",
  "is_active": true
}
```


