# readme todo
- [x] dev guide
- [ ] usage guide
- [x] link to how to release
- [x] future features
- [x] top summary section
- [x] features list
- [x] differences to dak.gg

# ER Builds Analysis Tool
Item stats display tool for [Eternal Return](https://store.steampowered.com/app/1049590/Eternal_Return/). Currently only supports windows.

ER Builds is based on [dak.gg](https://dak.gg/er/characters), but instead of presenting only the top 5 items per item category:

*(image from dak.gg)*

![](./doc/img/1.png)

ER Builds presents all the items per category, although it only supports Purple items (for now). This can be useful during build construction when you want to see more than just the top 5 items. Additionally, it presents more statistics than just Pick % and Average Rank.

*(image of ER Builds tool)*

![](./doc/img/2.png)

# Usage
[Usage Guide](./doc/usage-guide.md)

# Features
- **Webpage display**: Collected information is displayed on a UI page similar to dak.gg.
- **Multiple Sort Options**: Item statistics can be sorted by various values.
    - Number of Builds and Build %
    - Number of Likes
    - Average and Highest Win Rate
- **Offline Data Store**: ER Builds collects data from API and stores on disk, so accessing the API is not spammed. User can easily choose which characters they would like to retrieve data for, and can update on demand.

# Possible Future Features
- **Support for Yellow Items, Tac Skills, Augments**: Currently, the tool only collects information on Purple builds. It would be useful to know what the pick rates of Yellow Items, Tac Skills and Augments would be, but this requires additional collection of data.
- **UI For Selecting Characters for Data Collection**: Right now this feature is configured through a file, but a UI can make this better.
- **Item Stats Display**: The tool does not yet display the stats of an item on hovering over it like dak.gg.

# Development Information
[Development Notes](./doc/dev-notes.md)