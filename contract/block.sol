pragma solidity ^0.8.0;

import "./core/token/ERC721/extensions/ERC721URIStorage.sol";
import "./core/utils/Counters.sol";

contract PixelBlockItem is ERC721URIStorage {
    
    using Counters for Counters.Counter;
    
    Counters.Counter private _tokenIds;

    constructor() ERC721("PixelBlockItem", "PBI") {}
    
    function awardItem(address player, string memory tokenURI) public returns (uint256){
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();
        _mint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }
    function burned(uint256 tokenId) public {
        _burn(tokenId);
    }
}