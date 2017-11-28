package mptree

import (
	"github.com/goodgamecoin/ggcoin/common/rlp"
	//"github.com/rs/zerolog/log"
	"io"
)

func (n *node) EncodeRLP(w io.Writer) error {
	if err := rlp.Encode(w, n.path); err != nil {
		return err
	}
	children := []interface{}{}
	for i := byte(0); i < 255; i++ {
		if val := n.children[i]; len(val) > 0 {
			i := i
			children = append(children, &i)
			children = append(children, &val)
		}
	}
	if err := rlp.Encode(w, children); err != nil {
		return err
	}
	data := []interface{}{}
	for i := byte(0); i < 255; i++ {
		if val := n.data[i]; len(val) > 0 {
			i := i
			data = append(data, &i)
			data = append(data, &val)
		}
	}
	if err := rlp.Encode(w, data); err != nil {
		return err
	}
	return nil
}

func (n *node) DecodeRLP(s *rlp.Stream) error {
	path, err := s.Bytes()
	if err != nil {
		return err
	}
	children, err := readMap(s)
	if err != nil {
		return err
	}
	data, err := readMap(s)
	if err != nil {
		return err
	}
	*n = node{
		path:     path,
		children: children,
		data:     data,
	}
	return nil
}

func readMap(s *rlp.Stream) (map[byte][]byte, error) {
	_, err := s.List()
	if err != nil {
		return nil, err
	}
	m := map[byte][]byte{}
	for {
		index, err := s.Uint()
		if err != nil {
			if err == rlp.EOL {
				break
			}
			return nil, err
		}
		value, err := s.Bytes()
		if err != nil {
			return nil, err
		}
		m[byte(index)] = value
	}
	if err := s.ListEnd(); err != nil {
		return nil, err
	}
	return m, nil
}
